import file from 'k6/x/fs';

const root = '.';

export default function () {
    const jsFiles = file.walkMatch(root, '*.js');
    console.log(jsFiles)
}