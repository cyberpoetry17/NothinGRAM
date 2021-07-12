function validateJWS(tokenInfo) {
    let isExpired = false;
    let currentDate = new Date();    
    if (tokenInfo.exp*1000 < currentDate.getTime()) {
        isExpired = true;
    }

    return isExpired;
}

export default validateJWS;