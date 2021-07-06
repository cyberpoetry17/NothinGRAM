function validateJWS(tokenInfo) {
    var isExpired = false;
    var currentDate = new Date();
    if (tokenInfo.exp < currentDate.getTime()) {
        isExpired = true;
    }

    return isExpired;
}

export default validateJWS;