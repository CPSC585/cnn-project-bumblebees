from __future__ import print_function
import pandas as pd
import numpy as np
import csv
import sys
import argparse
import os
import logging
sys.path.insert(0, 'src')
from src.utils import create_paths, one_hot_maker



BASE_OUTPUT_DIR = "" #pwd


def build_parser():
    parser = argparse.ArgumentParser(description='generate the final label format for multi-label output')

    parser.add_argument('--label-path', type=str,
                        dest='label_path',
                        help='string path to the orignal label file',
                        required=True)

    parser.add_argument('--sort-by', type=str,
                        dest='sort_by',
                        help='row to sort the dataset',
                        required=True)

    parser.add_argument('--extract-from', type=str,
                        dest='extract_from',
                        help='row to extra labels',
                        required=True)

    parser.add_argument('--dest-path', type=str,
                        dest='dest_label_path',
                        help='string path to the destination of label csv',
                        required=True)

    parser.add_argument('--set-filename', type=str,
                        dest='filename',
                        help='filename for preprocessed labels csv',
                        required=True)

    parser.add_argument('--header', type=int,
                        dest='header', help='whether include first row',
                        metavar='HEADER', default=0)


    return parser



if __name__ == '__main__':
    parser = build_parser()
    options = parser.parse_args()
    print("Creating path for dest file...")
    create_paths(options)
    logging.basicConfig(filename=os.path.join(options.dest_label_path, 'preprocessing.log'), level=logging.DEBUG)
    logging.debug("Creating path for dest file...")
    """
    print(options.label_path)
    print(options.sort_by)
    print(options.extract_from)
    print(options.dest_label_path)
    print(options.filename)
    print(options.header)
    """

    print("generating final labels...")
    logging.debug("generating final labels...")
    one_hot_maker(options,logging)
    print("finished!")
    logging.info("finished!")





