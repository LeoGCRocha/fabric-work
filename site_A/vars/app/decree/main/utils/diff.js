'use strict'

/**
 * Process transactions of type read
 * @param {json} old - Old xml json to compare
 * @param {json} new - New xml json to compare
 */
const jsonDiff = require('json-diff')
function diffHist(old_, new_) {
    let diff = jsonDiff.diff(old_, new_)
    if (diff != undefined) {
        return true
    } else {
        return false
    }
}

module.exports = {diffHist}