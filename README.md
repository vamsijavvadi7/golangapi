# golangapi
login check:
https://api421.herokuapp.com/login/email/{mail}
Response:
{
"status": "True",
"role": "faculty"
}

FacultyDashboard ->Faculty Details(name,speciality..)
https://api421.herokuapp.com/fdashboard/details/{email}
Response:
{
"details": {
"name": "jyothiadabala",
"speciality": "surgeon"
}
}
FacultyDashboard ->Competency List
https://api421.herokuapp.com/fdashboard/competencydetails/{speciality}
Response:
{
"details": [
{
"competencyname": "competency1",
"competencyid": 4
},
{
"competencyname": "competency2",
"competencyid": 14
},
{
"competencyname": "competency3",
"competencyid": 24
},
{
"competencyname": "competency4",
"competencyid": 34
},
{
"competencyname": "competency5",
"competencyid": 44
},

]
}
FacultyDashboard ->List of Students and their percentages in a Competency
https://api421.herokuapp.com/fdashboard/competencydetails/speciality/{speciality}/competencyid/{competencyid}
Response:
[
{
"name": "kdajjnvuhdikauh",
"regno": "18pzqsqw1",
"self": 50,
"faculty": 50,
"competencynum": 4,
"competencyname": "competency1"
},
{
"name": "hgjugyjkuhjuyjgjkhu",
"regno": "18pa1a05b7",
"self": 71,
"faculty": 60,
"competencynum": 4,
"competencyname": "competency1"
},
{
"name": "ddsjugyjkuhjadfcjgjkhu",
"regno": "1oi2hw1owm1n2",
"self": 0,
"faculty": 0,
"competencynum": 4,
"competencyname": "competency1"
}]
Profile Details for all student or faculty based on their role
https://api421.herokuapp.com/profile/email/{email}
Response for faculty:
[
{
"name": "hkiuhkjkhkjkhvh",
"role": "faculty",
"phone": "87988767",
"email": "pramilacheruku001@gmail.com",
"batch": ""
}
]
Response for faculty:
[
{
"name": "hkiuhkjkhkjkhvh",
"role": "student",
"phone": "87988767",
"email": "pramilacheruku001@gmail.com",
"batch": "2019"
}
]
FacultyDashboard ->List of Students and their percentages in a Competency->Get Evaluations of a student in a competency
https://api421.herokuapp.com/competencyevaluations/competencyid/{competencyid}/studentid/{studentid}
Response:
[
{
"compentencyevaluationid": 4,
"patientop": "123",
"date": "08-01-2022",
"time": "09:43:46",
"self": 50,
"faculty": 50
},
{
"compentencyevaluationid": 24,
"patientop": "125",
"date": "08-01-2022",
"time": "09:44:13",
"self": 100,
"faculty": 100
},
{
"compentencyevaluationid": 54,
"patientop": "126",
"date": "08-01-2022",
"time": "12:43:22",
"self": 0,
"faculty": 0
}]
FacultyDashboard ->List of Students and their percentages in a Competency->Get Evaluations of a student in a competency
->Create a row in evaluation table with opnum(Post request):
https://api421.herokuapp.com/competencyevaluations/competencyid/{competencyid}/studentid/{studentid}/opnum
body:{"opnum":num,"fmail":""}
Response:
same as body

FacultyDashboard ->List of Students and their percentages in a Competency->Get Evaluations of a student in a competency
->Get Evaluation Form 
https://api421.herokuapp.com/competencyevaluations/competencyid/{competencyid}/competencyevaluationid/{competencyevaluationid}
Response:
{
"compevaluationid": 184,
"patientop": "128",
"date": "10-01-2022",
"time": "05:37",
"studentname": "hgjugyjkuhjuyjgjkhu",
"facultyname": "hkiuhkjkhkjkhvh",
"criteriadetails": [
{
"criteiaid": 4,
"criteriaqs": "how",
"option0": "not accurate",
"option1": "partially correct",
"option2": "perfect"
},
{
"criteiaid": 14,
"criteriaqs": "what",
"option0": "not accurate",
"option1": "partially correct",
"option2": "perfect"
},
{
"criteiaid": 24,
"criteriaqs": "why",
"option0": "not accurate",
"option1": "partially accurate",
"option2": "perfect"
},
{
"criteiaid": 34,
"criteriaqs": "when",
"option0": "not accurate",
"option1": "partially accurate",
"option2": "perfect"
}
]
}
FacultyDashboard ->List of Students and their percentages in a Competency->Get Evaluations of a student in a competency
->Post Evaluation Form (Post request)
https://api421.herokuapp.com/competencyevaluations/competencyid/{competencyid}/studentid/{studentid}
body:{[
{"criteriaid":id,
"score":1,
"matter":"dsds"
},{"criteriaid":id,
"score":1,
"matter":"dsds"
},{"criteriaid":id,
"score":1,
"matter":"dsds"
}]
}
Response:
same as body

Todo tasks->Faculty Meets->Get List of Students who want a one to one meet
https://api421.herokuapp.com/facultytodo/meet/{mail}
Response:
[
{
"studentname": "hgjugyjkuhjuyjgjkhu",
"competencyname": "competency1",
"studentid": "18pa1a05b7",
"CompetencyEvaluation_id": 184
}
]
Todo tasks->Faculty Reference Needed
https://api421.herokuapp.com/facultytodo/reference/{mail}
mailexample:pramilacheruku001@gmail.com
Response:
[
{
"studentname": "hgjugyjkuhjuyjgjkhu",
"competencyname": "competency1",
"studentid": "18pa1a05b7",
"criteriaqs": "what",
"CompetencyEvaluation_id": 184
},
{
"studentname": "hgjugyjkuhjuyjgjkhu",
"competencyname": "competency1",
"studentid": "18pa1a05b7",
"criteriaqs": "why",
"CompetencyEvaluation_id": 194
},
{
"studentname": "hgjugyjkuhjuyjgjkhu",
"competencyname": "competency1",
"studentid": "18pa1a05b7",
"criteriaqs": "why",
"CompetencyEvaluation_id": 184
}]
StudentDashboard ->Student Details(name,batch)
https://api421.herokuapp.com/studentdashboard/details/studentmail/{email}
Response:
{
"name": "kdajjnvuhdikauh",
"batch": "2018"
}
StudentDashboard ->Speciality List
https://api421.herokuapp.com/studentdashboard/specialities
Response:
{
"details": [
{
"specialityName": "surgeon",
"SpecialityId ": 1
},
{
"specialityName": "dentist",
"SpecialityId ": 2
},
{
"specialityName": "underteeth",
"SpecialityId ": 3
},
{
"specialityName": "upperteeth",
"SpecialityId ": 4
}
]
}
StudentDashboard ->>List Competency and student percentage in that competency
https://api421.herokuapp.com/studentdashboard/email/{mail}/speciality/{speciality}
Response:
[
{
"competencyid": 4,
"competencyname": "competency1",
"self": 50,
"faculty": 50
},
{
"competencyid": 14,
"competencyname": "competency2",
"self": 0,
"faculty": 0
},
{
"competencyid": 24,
"competencyname": "competency3",
"self": 0,
"faculty": 0
},
{
"competencyid": 34,
"competencyname": "competency4",
"self": 0,
"faculty": 0
},
{
"competencyid": 44,
"competencyname": "competency5",
"self": 0,
"faculty": 0
}]
