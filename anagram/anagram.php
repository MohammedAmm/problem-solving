<?php

//Check if it has only english char
function hasEnglishLetters(string $myString):bool
{
    if (preg_match('/[^A-Za]/', $myString))
    {
        return false;
    }
    return true;
}

//Validate string
function isValid(string $firstString, string $secondString):bool
{
    //Return False if they are not the same length
    if (strlen($firstString)!=strlen($secondString)) {
        return false;
    }
    //Check first string
    hasEnglishLetters($firstString);
    //Check second string
    hasEnglishLetters($secondString);
    //Return False
    return true;
}

function areAnagram(string $firstString, string $secondString):bool
{
    //Convert to small letters to make it easy to deal with.
    $firstString=strtolower($firstString);
    $secondString=strtolower($secondString);
    
    //Validation
    if(!isValid($firstString, $secondString)){
        return false;
    }

    //English letters 26 chars so init array with 26 chars
    $count=[];
    for ($i=0; $i < 26; $i++) { 
        $count[$i]=0;
    }

    //Loop through the two strings and increment in array if char in first string and decrement if char in second string
    //Get char position by reducing 97 dealing with small letters
    for ($i=0; $i < strlen($firstString); $i++) { 
        $count[ord($firstString[$i])-97]++; 
        $count[ord($secondString[$i])-97]--;
    }
    for ($i=0; $i < count($count); $i++) { 
        if ($count[$i]) {
            return false;
        }
    }
    return true; 
}

print(areAnagram('silent', 'liste1N'));