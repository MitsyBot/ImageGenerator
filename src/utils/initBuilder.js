const profile = require('../renderers/ProfileRender')
const canvas = require('./CanvasPrototypes')
const loadAllNeeded = async () => {
  profile.start()
  canvas.start()
}

module.exports = { loadAllNeeded }