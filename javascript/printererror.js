function printerError(s) {
   // your code
   let count = 0
   const wrongColor = ['n','o','p','q','r','s','t','u','v','w','x','y','z']
   for(let i = 0; i < s.length; i++) {
      if(wrongColor.includes(s[i])) count++
    }

    return `${count}/${s.length}`
}

console.log(printerError('aaaxbbbbyyhwawiwjjjwwm'))

