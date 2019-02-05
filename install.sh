TAG=$1;

if (( $EUID != 0 )); then
    echo "Please run command as root.";
    exit;
fi

DOWNLOADER="https://raw.githubusercontent.com/ml27299/lit-cli/master/godownloader.sh";
echo "yo";
eval "curl -sL -o- ${DOWNLOADER} | bash -s -- -b /usr/local/bin $TAG"