'use strict'

const { json } = require('express')
const express = require('express')
const mini = require('./main')
const errorHandler = require('./utils/errorHandler')

const port = 8080
const host = '0.0.0.0'

const app = express()
app.use(json())

app.post('/write', async (req, res, next) => {
    let result
    try {
        await mini.start(req.body.data)
    } catch (error) {
        return next(error)
    }
    const returnJson = {
        message: "Successfully executed application",
        data: result,
        success: true
    }
    res.status(200).json(returnJson)
    return next()
})
app.post('/read', async (req, res, next) => {
    let result
    try {
        result = await mini.start(req.body.data)
    }
    catch (error) {
        return next(error)
    }
    const returnJson = {
        message: "Successfully executed application",
        success: true,
        result
    }
    res.status(200).json(returnJson)
    return next()
})
app.post('/update', async(req, res, next) => {
    let result 
    try {
        result = await mini.start(req.body.data)
    } catch (error) {
        return next(error)
    }
    const returnJson = {
        message: "Successfully executed application",
        success: true,
        result: result
    }
    res.status(200).json(returnJson)
    return next()
})
app.use(errorHandler)
app.listen(port, host)