isSubsequence("aaaaaa", "bbaaaa")

function isSubsequence(s, t) {
    let sPointer = 0;
    let tPointer = 0;

    while (sPointer < s.length && tPointer < t.length) {
        if (s[sPointer] === t[tPointer]) {
            sPointer++;
        }
        tPointer++;
    }

    return sPointer === s.length;
}
 

/**
 * 
        const subString = t.substring(lastIndex);

        let i = subString.indexOf(c); 

        

        if (i == -1) {
            return false;
        } 

        lastIndex = i;
         
        console.log(c, i , subString, lastIndex);
 */