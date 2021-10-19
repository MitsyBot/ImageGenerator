const { buildProfileImage } = require("../renderers/ProfileRender")
const { buildShipImage } = require("../renderers/ShipRender")

const renderShip = async (req, res) => {
  const { linkOne, linkTwo, shipValue } = req.body
  const result = await buildShipImage(linkOne, linkTwo, shipValue)
  res.send(result.toJSON())
}

const renderProfile = async (req, res) => {
  const { user, marry, usageCommands, i18n } = req.body
  const result = await buildProfileImage(user, marry, usageCommands, i18n)
  res.send(result.toJSON())
}


module.exports = {
  renderShip,
  renderProfile,
}