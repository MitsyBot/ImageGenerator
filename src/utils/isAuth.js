const isAuthorized = (req, res, next) => {

    const auth = req.headers.authorization
  
    if (!auth || auth !== process.env.API_TOKEN) {
      return res.status(401).send({ message: 'Apenas a mitsy pode acessar isso!' });
    }
  
    return next();
  };
  
  module.exports = { isAuthorized }