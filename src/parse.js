function Person(name, sex, birth, death, buried, occupation, education, religion) {
  this.name       = name;
  this.sex        = sex;
  this.birthDate  = birth.date;
  this.birthPlace = birth.place;
  this.deathDate  = death.date;
  this.deathPlace = death.place;
  this.buried     = buried;
  this.occupation = occupation;
  this.education  = education;
  this.religion   = religion;
}

function openFile(file) {
  var xhr = new XMLHttpRequest();
  xhr.onload = readGEDCOM;
  xhr.open('GET', file);
  xhr.send();
}

function readGEDCOM() {
  var unparsed = this.responseText.split("\r");
  for (var i = 0; i < numRecordsAtZero(unparsed); i++) {

  }
}

function numRecordsAtZero(arr) {
  var numAtZero = 0;
  for (var i = 0; i < arr.length; i++) {
    if (arr[i][0] === '0') numAtZero++;
  }
  return numAtZero;
}
openFile('http://0.0.0.0:8000/src/sample.ged');
