#!/bin/bash
#
# Call this script with an email address (valid or not),
# to create a test certificate and key for use with a TLS connection.

if [ -z $1 ]; then
	echo "Usage: makecert.sh joe@random.com";
	exit 1;
fi

echo "This script generates 1024 bit RSA key and certificate files.";
echo "These files are strictly for testing purposes.";
echo "Do NOT use them in production services.";

mkdir -p certs;
rm certs/*;

openssl req -new -nodes -x509 -out certs/cert.pem -keyout certs/key.pem -days 3650 \
	-subj "/C=DE/ST=NRW/L=Earth/O=Random Company/OU=IT/CN=www.random.com/emailAddress=$1";

exit $?;
