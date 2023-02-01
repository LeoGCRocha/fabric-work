'use strict';

const CustomError = require('./CustomError')

// eslint-disable-next-line no-unused-vars
function errorHandler(err, req, res, next) {
    if (!(err instanceof CustomError)) {
        console.log(err)
        res.status(500).send({
            message: "Ocorreu um erro não identificado",
            success: false
        })
        return
    }

    err.log()
    res.status(err.httpStatus).send({
        message: err.message,
        success: false
    })
}

module.exports = errorHandler;