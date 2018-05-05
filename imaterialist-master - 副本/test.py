from __future__ import print_function
import pandas as pd
import numpy as np
import csv
import sys
import argparse
import os
sys.path.insert(0, 'src')
from src.utils import create_paths, one_hot_maker


import csv
import numpy as np

if __name__ == '__main__':
    for i in range(10):
        filename = "/temworkspace/imaterialist-master/testdata/train.csv"
        csvfile = open(filename, "a+", newline='')
        # csvfile = open(pathAttributes.face_features_data,"w",newline='')  batch mode
        writer = csv.writer(csvfile)
        k = [[1, 2, 3]]
        writer.writerows(k)
        csvfile.close()






