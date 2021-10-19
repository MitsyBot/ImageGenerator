const { Router } = require('express')
const controller = require('./controllers/controllers')

const router = Router()

router.get('/ship', controller.renderShip)
router.get('/profile', controller.renderProfile)


module.exports = router;