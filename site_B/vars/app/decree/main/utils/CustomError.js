'use strict'

const errors = require('./errors.json')

class CustomError extends Error {
    #name
    #httpStatus
    
    constructor(name, context, ...args) {
        super(CustomError.#getMessage(name, context, ...args));
        this.#name = name
        this.#httpStatus = errors[name].httpStatus
    }

    get name() { return this.#name; }
    get httpStatus() { return this.#httpStatus; }

    static #format = (str, ...args) => {
        args.forEach(function (arg) {
            str = str.replace("%s", arg);
        })
        return str
    };
    
    static #getMessage = (name, context, ...args) => {
        let error = errors[name]
        
        let message
        if (error === undefined) {
            message = context + " | " + this.#format("Error: %s | Erro não encontrado", name)
            return message
        }
    
        message = context + " | " + this.#format(error.templateMessage, ...args)
    
        return message
    }

    getLog = () => {
        return this.stack
    }

    log = () => {
        console.log(this.stack)
    }
}

module.exports = CustomError