FILE=go.mod

if [ -f "$FILE" ]; 
then
    echo "starting the program..."
    go run github.com/eyupfatihersoy/app-tryout-1
else
    echo "$FILE does not exist."
fi

echo "finished executing."