import socket
# import qrcode
import argparse
import qrcode_terminal

def get_local_ip():
    """Get the local IP address of the current machine."""
    hostname = socket.gethostname()
    local_ip = socket.gethostbyname(hostname)
    return local_ip

def generate_qr_code(data):
    """Generate and display a QR code for the given data."""
    # qr = qrcode.QRCode(
    #     version=1,
    #     error_correction=qrcode.constants.ERROR_CORRECT_L,
    #     box_size=10,
    #     border=4,
        
    # )
    # qr.add_data(data)
    # qr.make(fit=True)
    # qr_terminal = qrcode_terminal.draw(str(qr))
    # qr_terminal.show()

    qrcode_terminal.draw(data)

def main(port=80, use_https=False):
    local_ip = get_local_ip()
    protocol = "https" if use_https else "http"
    address = f"{protocol}://{local_ip}:{port}"
    print(f"Generated address: {address}")
    generate_qr_code(address)

if __name__ == "__main__":
    parser = argparse.ArgumentParser(description="Generate a QR code for the local IP address.")
    parser.add_argument("--port", type=int, default=80, help="Port to append to the local address. Default is 80.")
    parser.add_argument("--https", action="store_true", help="Use HTTPS instead of HTTP.")
    args = parser.parse_args()
    
    main(port=args.port, use_https=args.https)
