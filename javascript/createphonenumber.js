function createPhoneNumber(phoneNumber) {
    let phone = ''
    phoneNumber.forEach((e,i) => {
        switch(i) {
            case 0 : 
                phone += '('
                phone += e
            break
            case 2: 
                phone +=e
                phone += ") "
                break
            case 5 : 
                phone += e
                phone += "-"
                break
            default :
            phone += e
        }        
    });
    return phone
}


createPhoneNumber([1, 2, 3, 4, 5, 6, 7, 8, 9, 0])