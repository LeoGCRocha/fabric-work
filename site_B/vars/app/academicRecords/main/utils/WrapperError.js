'use strict'

const CustomError = require('./CustomError')

class WrapperError extends CustomError {
    #wrappedError
    
    constructor(name, context, wrappedError, ...args) {
        super(name, context, ...args);
        this.#wrappedError = wrappedError
    }

    get wrappedError() { return this.#wrappedError; }

    getLog = () => {
        return `${this.name}: ${this.message}\n${JSON.stringify(this.#wrappedError)}`
    }

    log = () => {
        console.log(this.stack)
        console.log(this.#wrappedError)
    }
}

module.exports = WrapperError