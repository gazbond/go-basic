const fse = require('fs-extra');

fse.copySync("./node_modules/bootstrap/dist", "./static/assets/dist", { overwrite: true })