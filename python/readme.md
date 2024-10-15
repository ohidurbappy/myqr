### Usage:
Run the script with the default settings:
```bash


python script_name.py
```


### Specify a custom port:


```bash
python script_name.py --port 8080
```


### Use HTTPS:


```bash
python script_name.py --https
```

### Use both custom port and HTTPS:


```bash
python script_name.py --port 8080 --https
```

### Description:

This script will generate a QR code for the local address with the specified port and protocol, then display it in the terminal.


### On Mac

sudo chmod +x ./myqr
sudo mv ./myqr /usr/local/bin/myqr