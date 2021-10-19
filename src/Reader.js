const fs = require('fs')

const BuferedImage = (image) => {
  return fs.readFileSync(`src/assets/images/${image}.png`)
}

const getBalance = () => BuferedImage('balance');
const getBravery = () => BuferedImage('bravery');
const getBrilliance = () => BuferedImage('brilliance');


module.exports = {
  getBalance,
  getBravery,
  getBrilliance,
}