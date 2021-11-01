FILE=go.mod

if [ -f "$FILE" ]; 
then
    echo "starting the program..."
    echo "..."
    go run github.com/eyupfatihersoy/app-tryout-1
else
    echo "$FILE does not exist."
    echo "configuring files..."
    go mod init
    echo "..."
    go run github.com/eyupfatihersoy/app-tryout-1
fi

echo "finished executing."