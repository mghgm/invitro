import pandas as pd
import logging as log

log.basicConfig(
        level=log.INFO,
        format='(%(asctime)s) Trace synthesizer -- [%(levelname)s] %(message)s'
    )

def load_data(path):
    return pd.read_csv(path)

def save_data(data, path):
    failed = True
    while failed:
        try:
            data.to_csv(path, index=False)
            failed = False
        except:
            log.warn(f'Failed to save {path}')
            failed = True 
