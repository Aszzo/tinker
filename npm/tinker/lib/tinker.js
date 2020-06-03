#!/usr/bin/env node
const program = require('commander');
const pkgConfig = require('../package');
const child_process = require('child_process');
const path = require('path');
let  flags = [];
program
    .version(pkgConfig.version, '-v, --version')
    .usage('[options]')
    .requiredOption('-i, --input <value>', '目标文件所在的目录|目标文件路径')
    .option('-q, --quality <number>', '压缩图片的质量(0-100)', 80)
;
program.parse(process.argv);

if (program.input) {
    flags = [...flags, ...['-i', program.input]]
}
flags = [...flags, ...['-q', program.quality]];

const tinker = path.join(__dirname, '..', 'bin', 'tinker');

const child = child_process.spawn(tinker, flags);

child.stdout.on('data', data => {
    console.log(data.toString())
})

child.stderr.on('data', (data) => {
    console.error(`err: ${data}`);
});