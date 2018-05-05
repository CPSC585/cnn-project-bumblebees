from __future__ import print_function
import os
import glob
import yaml
from shutil import copyfile
import importlib
import pandas as pd
import numpy as np
import csv
from sklearn.preprocessing import OneHotEncoder

def exists(p, msg):
    assert os.path.exists(p), msg


def importmod(module, msg):
    try:
        return importlib.import_module("{}.{}".format(module, "network"))
    except ImportError:
        print (msg)

def read_train_labels_file(opts, log):
    if os.path.exists(opts.train_target_path):  #fixed the header bug
        log.debug("Loading labels...")
        labels = pd.read_csv(opts.train_target_path, header=None)
        opts.train_labels = np.array(labels[:])
    else:
        log.warn("path does not exist.")
        opts.train_labels = np.array([])

    return opts

def read_val_labels_file(opts, log):
    if os.path.exists(opts.val_target_path):
        log.debug("Loading labels...")
        labels = pd.read_csv(opts.val_target_path, header=None)   #fixed the header bug
        opts.val_labels = np.array(labels[:])
    else:
        log.warn("path does not exist.")
        opts.val_labels = np.array([])

    return opts

def create_paths(opts):
    try:
        if not os.path.exists(opts.log_path):
            os.makedirs(opts.log_path)
    except:
        print("no log_path")
    try:
        if not os.path.exists(opts.dest_label_path):
            os.makedirs(opts.dest_label_path)
    except:
        print("no label_path")

    try:
        if not os.path.exists(opts.tf_log_path):
            os.makedirs(opts.tf_log_path)
    except:
        print("no tf_log_path")

def _read_aug_params(path):
    with open(path, 'r') as stream:
        data_loaded = yaml.load(stream)
    return data_loaded


def read_val_aug_params(opts, log):
    path = os.path.join(opts.model, 'augment.yml')
    if os.path.exists(path) and opts.aug_val:
        copyfile(path, os.path.join(opts.output_path, os.path.basename(path)))
        log.debug("Loading Validation Augmentation Parameters...")
        opts.val_aug_params = _read_aug_params(path)
        #for key, value in opts.val_aug_params.iteritems(): python2
        for key, value in opts.val_aug_params.items(): #python3
            log.info("Validation augmentation param -- {}: {}".format(key, value))
    else:
        opts.val_aug_params = {}
        log.warn("NO Validation Augmentation.")
    return opts


def read_train_aug_params(opts, log):
    path = os.path.join(opts.model, 'augment.yml')
    if os.path.exists(path) and opts.aug_train:
        copyfile(path, os.path.join(opts.output_path, os.path.basename(path)))
        log.debug("Loading Train Augmentation Parameters...")
        opts.train_aug_params = _read_aug_params(path)
        #for key, value in opts.train_aug_params.iteritems(): python2
        for key, value in opts.train_aug_params.items(): #python3
            log.info("Train augmentation param -- {}: {}".format(key, value))
    else:
        opts.train_aug_params = {}
        log.warn("NO Train Augmentation.")
    return opts


def read_unfreeze_layers(opts, log):
    path = os.path.join(opts.model, 'unfreeze_layers')
    if opts.freeze_layers and os.path.exists(path):
        copyfile(path, os.path.join(opts.output_path, os.path.basename(path)))
        with open(path, 'r') as f:
            opts.unfreeze_layers = f.readlines()
            opts.unfreeze_layers = list(map(lambda s: s.strip(), opts.unfreeze_layers))
    elif opts.freeze_layers and not os.path.exists(path):
        log.warn("Specified freezing layers but no file with information about freeze layers found! All Layers are Trainable.")
    return opts


def setup_paths(opts, base_dir, log):
    # TODAY.strftime("%d-%b-%Y")
    name = "opt-{}_s-{}".format(opts.optimizer, opts.seed)
    if opts.aug_train:
        name += "_aug_train"
    if opts.aug_val:
        name += "_aug_val"
    if opts.pretrained:
        name += "_pretrained"
    if opts.freeze_layers and os.path.exists(os.path.join(opts.model, 'unfreeze_layers')):
        name += "_partly_frozen"
    opts.output_path = os.path.join(base_dir, opts.model, name)
    opts.log_path = os.path.join(opts.output_path, 'logs')
    opts.log_file = os.path.join(opts.log_path, 'train.csv')
    opts.tf_log_path = os.path.join(opts.output_path, 'tf_logs')
    create_paths(opts)
    return opts



def check_opts(opts, log):
    assert opts.epochs > 0
    log.info("Number of Epochs: {}".format(opts.epochs))
    assert opts.batch_size > 0
    log.info("Batch Size: {}".format(opts.batch_size))
    if not os.path.isdir(opts.train_path):
        log.error("Train path {} does not exist!".format(opts.train_path))
        raise IOError("Train path {} does not exist!".format(opts.train_path))
    else:
        log.info("Training Data: {}".format(opts.train_path))
    if not os.path.isdir(opts.val_path):
        log.error("Train path {} does not exist!".format(opts.val_path))
        raise IOError("Train path {} does not exist!".format(opts.val_path))
    else:
        log.info("Validation Data: {}".format(opts.val_path))


def import_model(opts):
    opts.model = importmod(opts.model, "Model {} does not exist!".format(opts.model))


def get_saved_models(opts):
    files = []
    if os.path.exists(opts.output_path):
        files = glob.glob(os.path.join(opts.output_path, '*.hdf5'))
        files.sort(key=lambda x: float(os.path.basename(x).split("-")[1].replace(".hdf5", "").replace("val_loss=", "")))
        return files

def one_hot_maker(options, logging):
    labels = pd.read_csv(options.label_path, header=options.header)
    count = labels.shape[0]
    batch_size = 3000
    classes = []
    for i in labels[options.extract_from]:
        k = i.split(' ')
        for j in k:
            classes.append(j)
    classes1 = np.array(classes)
    unique = np.unique(classes1)
    unique = unique.reshape(-1, 1)
    encoder = OneHotEncoder()
    encoder.fit_transform(unique)
    s = labels.sort_values(by=options.sort_by)
    for i in range(count//batch_size+1):
        print("{}-{}".format(i*batch_size,(i+1)*batch_size))
        logging.info("{}-{}".format(i * batch_size, (i + 1) * batch_size))
        """test code
        if i == 2:
            break
        """
        if i != count//batch_size:
            v_labels = s[options.extract_from][i*batch_size:(i+1)*batch_size]
        else:
            v_labels = s[options.extract_from][i * batch_size:count]
        v_labels_list = v_labels.tolist()
        output_list = []
        for i in v_labels_list:
            v_list = i.split(' ')
            output_list.append(label_output(v_list, encoder))
        print("writing final labels to file...")
        logging.debug("writing final labels to file...")
        msg = write_label_file(options, output_list, logging)
        print(msg)
        logging.info(msg)

def label_output(k, encoder):
    #preprocessed_labels = []
    y = 0
    for i in k:
        curr_output = encoder.transform(i)
        y = y + curr_output
        #preprocessed_labels.append(y.toarray())
    return y.toarray().reshape(-1)

def write_label_file(options, output, logging):
    meg = ""
    try:
        filename = os.path.join(options.dest_label_path, options.filename)
        print("storing in {}".format(filename))
        logging.debug("storing in {}".format(filename))
        csvfile = open(filename, "ab")  #be careful about a/w , python3 uses newline='', python2 uses b
        #print("can't open")
        # csvfile = open(pathAttributes.face_features_data,"w",newline='')  batch mode
        writer = csv.writer(csvfile, delimiter=',')
        #print("can't write")
        writer.writerows(output)
        #print("can't input")
        csvfile.close()  # 14,49,53,61,105,106,
        #print("can't close")
        msg = "succeed!"
    except:
        msg = "error!"
    return msg