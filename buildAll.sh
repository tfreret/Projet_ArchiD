source_dir=./cmd
build_dir=./exe

if [ ! -d "$build_dir" ]; then
    mkdir "$build_dir"
fi

if [ "$(ls -A "$build_dir")" ]; then
    rm -r "${build_dir:?}/"*
fi

files=("sensors/pressureSensor" "sensors/tempSensor" "sensors/windSensor" "fileRecorder" "databaseRecorder" "api" "alertManager")

for file in "${files[@]}"; do
    echo "Building $file"
    go build -o "$build_dir/$file" "$source_dir/$file"
done

echo "Build completed."