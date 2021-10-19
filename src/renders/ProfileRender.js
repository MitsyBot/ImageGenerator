const CanvasImport = require('canvas')

const buildProfileImage = async (user, marry, usageCommands, i18n) => {
  const canvas = CanvasImport.createCanvas(1080, 720);
  const ctx = canvas.getContext('2d');

  const baseColor = '#f895fb';
  ctx.fillStyle = baseColor;
  ctx.fillRect(0, 0, canvas.width, canvas.height);

  ctx.textAlign = 'center';
  ctx.fillText('Opa', 980, 290);

  return canvas.toBuffer();
}

module.exports = { buildProfileImage, start }