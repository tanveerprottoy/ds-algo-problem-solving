"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.General = void 0;
var General = /** @class */ (function () {
    function General() {
    }
    General.prototype.reverse = function (str) {
        var res = "";
        for (var i = str.length - 1; i >= 0; i--) {
            res += str[i];
        }
        return res;
    };
    return General;
}());
exports.General = General;
