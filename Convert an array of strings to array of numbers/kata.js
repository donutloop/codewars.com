function toNumberArray(stringarray){

  var num = []
  for (i = 0; i < stringarray.length; i++) {

      if (isInt(stringarray[i])) {
          num.push(parseInt(stringarray[i]))
      } else {
          num.push(parseFloat(stringarray[i]))
      }
  }
  return num
}

function isInt(n) {
   return n % 1 === 0;
}