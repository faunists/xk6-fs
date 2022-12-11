# xk6-fs
[k6](https://github.com/grafana/k6) extension for writing files, implemented using the
[xk6](https://github.com/grafana/xk6) system.

## Build
```shell
xk6 build v0.41.0 --with github.com/alexrios/xk6-fs@latest
```

## Example
```javascript
import file from 'k6/x/fs';

const root = '.';

export default function () {
    file.walkMatch(root, '*.js');
}
```

## Run sample script
```shell
./k6 run sample-script.js
```