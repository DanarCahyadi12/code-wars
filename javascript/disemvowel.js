function disemvowel(str) {
    const a = ['a','i','u','e','o']
    let result = ''
    str.split('').forEach(alp => {
        if(!a.includes(alp.toLowerCase())) result += alp
    })
    return result
}


console.log(disemvowel( "This website is for losers LOL!"))
