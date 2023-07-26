function accum(s) {
    let output = ''
	for(let i = 0; i <= s.length-1; i++) {
        for(let j = 0; j <= i; j++) {
            if(j === 0) output += s[i].toUpperCase()
            else output += s[i].toLowerCase()
        }
        output += i !== s.length-1 ? "-" : ""
    }
    return output
}

console.log(accum('cwAt'))