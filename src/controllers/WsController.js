const { buildProfileImage } = require("../renderers/ProfileRender")
const { buildShipImage } = require("../renderers/ShipRender")


const WebSocketController = async (socket, rawRequest) => {
    const { type, data, id } = JSON.parse(rawRequest.toString());
  
    switch (type) {
      case 'ship': {
        const { linkOne, linkTwo, shipValue } = data
        const result = await buildShipImage(linkOne, linkTwo, shipValue)
        socket.send(JSON.stringify({ id, res: result.toJSON() }))
        break;
      }
      case 'profile': {
        const { user, marry, usageCommands, i18n } = data
        const result = await buildProfileImage(user, marry, usageCommands, i18n)
        socket.send(JSON.stringify({ id, res: result.toJSON() }))
        break;
      }
    }
  };
  
  module.exports = WebSocketController