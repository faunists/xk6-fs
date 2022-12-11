# xk6-fs
[k6](https://github.com/grafana/k6) extension for dealing with the file system, implemented using the
[xk6](https://github.com/grafana/xk6) system.

## Build
```shell
xk6 build --with github.com/faunists/xk6-fs@latest
```

## Example
```javascript
import file from 'k6/x/fs';

const root = 'samples';

export default function () {
    const jsFiles = file.walkMatch(root, '*.js');
    console.log(jsFiles)
}
```

## Run sample script
```shell
make run-sample
```
then you should see something like:\
```log
INFO[0000] ["samples/falcon.js","samples/tiger.js"]      source=console
```
