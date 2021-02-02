function openNav() {
    document.getElementById("mySidebar").style.width = "250px";
    document.getElementById("main").style.marginRight = "250px";
}

function closeNav() {
    document.getElementById("mySidebar").style.width = "0px";
    document.getElementById("main").style.marginRight = "0px";
}

// Stack Overflow Answer
function toggleNav() {
    var sidenav = document.getElementById("mySidebar");
    var main = document.getElementById("main");
    sidenav.style.width = sidenav.style.width === "250px" ? '0' : '250px';
    main.style.marginRight = main.style.marginRight === "250px" ? '0' : '250px';
}
async function retrieveStudentData(id) {
    const response = await fetch("/transaction", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: id,
    })
    return response.json()
};

// how to use data ^^^^^
// data = {'student_id': 123124, 'name': asdasif, 'credits', } data.student_id, data.name, data.credits
// functions courtsey of Hammad for getStudent and drop down menu

async function retrieveCourseData(id) {
    var course;
    const response = await fetch("/courses", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: id,
    });
    return response.json();
};

async function retrieveRequirementData(id) {
    var req;
    const response = await fetch("/requirements", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: id,
    });
    return response.json();
};

async function retrieveCoursesTaken(id) {
    var req;
    const response = await fetch("/courses_taken", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: id,
    });
    return response.json();
};

async function retrieveAllCourses() {
    var req;
    const response = await fetch("/all_courses", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: 123,
    });
    return response.json();
};

async function postTest() {
    const s = await retrieveStudentData(23848083);
    console.log(s);
    const c = await retrieveCourseData(17);
    console.log(c.credits);
    const r = await retrieveRequirementData(10);
    console.log(r);
    const ct = await retrieveCoursesTaken(23848083);
    console.log(ct);
    const ca = await retrieveAllCourses();
    console.log(ca);
}

async function displayStudentData(id) {
    s = await retrieveStudentData(id);
    idf = document.getElementById('id');
    ids = document.getElementById('sid');
    ids.innerHTML = s.student_id;
    idf.setAttribute('value', s.student_id);
    n = document.getElementById('name');
    n2 = document.getElementById('sname');
    n3 = document.getElementById('studentname');
    n3.value = s.name;
    n2.innerHTML = s.name;
    n.setAttribute('value', s.name);
    maj = document.getElementById('major');
    maj.setAttribute('value', 'Computer Science');
    gp = document.getElementById('sgpa');
    gp.innerHTML = s.overall_gpa;

    cr = document.getElementById('cred');
    cr.innerHTML = s.credits;

    cr2 = document.getElementById('credn');
    cr2.innerHTML = 120 - parseInt(s.credits);
}

async function courseTaken(c_id, s_id) {
    var courses = await retrieveCoursesTaken(s_id);
    //console.log(courses);
    //console.log(typeof courses[i].course_id);
    //console.log(typeof c_id);
    for (i = 0; i < courses.length; i++) {
        //console.log(courses[i].course_id.trim() + "|||||" + c_id.trim());
        if (courses[i].course_id.trim() === c_id.trim()) {
            try {
                if (courses[i].in_progress.trim() === '1')
                    return "t";
                else
                    return "p";
            }
            catch(err){ return "p"; }
        }
    }
    return "n";
}

async function retrievePrerequisites(course_id) {
    var c = await retrieveCourseData(parseInt(course_id));
    var courses = [];
    if (typeof c.prereqs == 'undefined')
        return courses;
    var prqs = c.prereqs.split(',');
    //console.log(prqs);
    for (i = 0; i < prqs.length; i++)
    {
        if (prqs[i] == -1 || prqs[i].length > 10)
            continue;
        var c2 = await retrieveCourseData(parseInt(prqs[i].trim()));
        courses.push(c2);
    }
    return courses;
}

async function addButtonsByCourse(course, id, s_id) {
    var req = document.getElementById(id);
    if (req == null)
    {
        console.log("ERROR" + id);
        return 'no';
    }
    var color;
    var taken = "#5CAD96"
    var taking = "#9FC3F9"
    var ntaken = "#E77478"
    var t = await courseTaken(course.course_id.trim(), s_id);
    //console.log(t);
    if (t == 't')
        color = taking;
    else if (t == 'p')
        color = taken;
    else
        color = ntaken;
    //console.log(color);
    var btn = document.createElement('div');
    btn.className = "dropdown";
    var btn2 = document.createElement('button');
    btn2.className ='dropbtn';
    btn2.style.backgroundColor = color;
    btn2.style.color = "black";
    btn2.innerHTML = course.department + " " + course.course_number;
    btn2.addEventListener('click', function() {
        myFunction(course.name + id)
    });
    var btn3 = document.createElement('div');
    btn3.id = "myDropdown" + course.name + id;
    btn3.className = "dropdown-content";
    var btn4 = document.createElement('a');
    // btn4.href ='#';
    btn4.innerHTML = course.name + " (" + course.credits + " CR)";
    btn4.style.backgroundColor = color;

    var prereqs = await retrievePrerequisites(course.course_id);
    var prestr = "";
    //console.log(prereqs);
    for (i = 0; i < prereqs.length; i++) {
        prestr += prereqs[i].department + " " + prereqs[i].course_number;
        if (i != prereqs.length-1)
            prestr += ', ';
    }
    var btn5 = document.createElement('p');
    btn5.innerHTML = course.description;
    var btn6 = document.createElement('p');
    if (prestr != "")
        btn6.innerHTML = 'Pre-requisites: ' + prestr;
    btn4.appendChild(btn6)
    btn4.appendChild(btn5);
    btn3.appendChild(btn4);
    btn2.appendChild(btn3);
    btn.appendChild(btn2);
    req.appendChild(btn);
    if (color == ntaken)
        return 'no';
    if (color == taking)
        return 't';
    //console.log(req);
}

function setRequirementFulfilled(req_id) {
    str = "sn" + req_id.toString();
    console.log(str);
    var ele = document.getElementById(str);
    ele.style.backgroundColor = "#5CAD96"
    ele.style.color = "black"
    // ele.style.whiteSpace = "nowrap"
    ele.innerHTML = 'Requirement Fulfilled!'
}

function setRequirementTaking(req_id, cred) {
    str = "sn" + req_id.toString();
    var ele = document.getElementById(str);
    if (ele === null)
        return;
    ele.style.backgroundColor = "#9FC3F9";
    ele.style.color = "#000000";
    if (cred <= 0)
        ele.innerHTML = 'In progress...';
    else
        ele.innerHTML = 'In progress. You need ' + cred + ' more credits upon completion';
}
function setRequirementUnFulfilled(req_id) {
    str = "sn" + req_id.toString();
    var ele = document.getElementById(str);
    if (ele === null)
        return;
    ele.style.backgroundColor = "#E77478";
    ele.style.color = "black";
    ele.innerHTML = 'Still Needed'
}

async function displayCSbuttons(s_id) {
    courses3 = await retrieveAllCourses();
    for (l = 0; l < courses3.length; l++)
    {
        var s = await addButtonsByCourse(courses3[l], "c" + courses3[l].course_id.toString(), s_id);
        if (s == 'no')
            setRequirementUnFulfilled("c" + courses3[l].course_id.toString());
        else if (s == 't')
            setRequirementTaking("c" + courses3[l].course_id.toString(), 0);
        else
            setRequirementFulfilled("c" + courses3[l].course_id.toString());
    }
}

async function displayAllButtons(s_id) {
    courses2 = await retrieveAllCourses();
    //console.log(courses2
    for (k = 2; k <= 16; k++)
    {
        if (k == 12)
            continue;
        req = await retrieveRequirementData(k);
        credits = req.credits_required;
        takingCredits = 0;
        //console.log(credits);
        for (j = 0; j < courses2.length; j++)
        {
            //console.log(courses2[j]);
            if (courses2[j].requirement_fulfilled.length > 2)
            {
                if (courses2[j].requirement_fulfilled.split(',').includes(k.toString()))
                {
                    check = await addButtonsByCourse(courses2[j], k.toString(), s_id);
                    var s = await courseTaken(courses2[j].course_id, s_id);
                    if (s == 'p')
                        credits -= courses2[j].credits;
                    if (s == 't')
                        takingCredits += courses2[j].credits;
                }
            }
            else if (courses2[j].requirement_fulfilled == k.toString())
            {
                await addButtonsByCourse(courses2[j], k.toString(), s_id);
                var s = await courseTaken(courses2[j].course_id, s_id);
                if (s == 'p')
                    credits -= courses2[j].credits;
                if (s == 't')
                    takingCredits += courses2[j].credits;
            }
            if (credits <= 0)
            {
                var itaken = await setRequirementFulfilled(k.toString());
            }
            else if (takingCredits != 0)
            {
                setRequirementTaking(k.toString(), credits-takingCredits);
            }
        }
    }
}

async function cleanReqs() {
    for (i = 2; i <= 17; i++)
    {
        if (i == 12)
            continue;
        await setRequirementUnFulfilled(i);
        str = "";
        if (i == 17 || i == 10)
            str += "c";
        str += i.toString();
        requ = document.getElementById(str);
        if (requ !== null)
        {
            while (requ.firstChild)
                requ.removeChild(requ.firstChild);
        }
    }

    for (i = 1; i < 50; i++)
    {
        setRequirementUnFulfilled("c" + i.toString());
        requ = document.getElementById("c" + i.toString());
        if (requ !== null)
        {
            while (requ.firstChild)
                requ.removeChild(requ.firstChild);
        }
    }
}

async function initialize()
{
    var id = document.getElementById("sid2").value;
    if (id == "")
        id = 23848083
    console.log(id);
    document.getElementById("sid2").value = id;
    await cleanReqs();
    //c = await retrieveCourseData(17);
    //c2 = await retrieveCourseData(12);
    await displayStudentData(id);
    await displayAllButtons(id);
    await displayCSbuttons(id);
}

function formRefreshStopper() {
    var form = document.getElementById("sidform");
    function handleForm(event) { event.preventDefault(); } 
    form.addEventListener('submit', handleForm);
}

function getStudent() {
    var student = document.getElementById("emplid").value;
    document.getElementById("emplid").innerHTML = student;
    document.querySelector("body > table.toplevel.table_default > tbody > tr:nth-child(3) > td.block_n2_and_content > table > tbody > tr:nth-child(2) > td.block_content_outer > table > tbody > tr > td > table:nth-child(10) > tbody > tr:nth-child(3) > td > a");
}
function checkVisible(elm) {
    var rect = elm.getBoundingClientRect();
    console.log(rect);
    var viewHeight = Math.max(document.documentElement.clientHeight, window.innerHeight);
    return !(rect.bottom < 0 || rect.top - viewHeight >= 0);
}

function myFunction(id) {
    str = "myDropdown" + id;
    //console.log(str);
    var ele = document.getElementById(str)
    if (!checkVisible(ele))
    {
        console.log('hi');
        ele.classList.style.bottom = '500px';
    }
    ele.classList.toggle("show");
}

// Close the dropdown if the user clicks outside of it
window.onclick = function (event) {
    if (!event.target.matches('.dropbtn')) {
        var dropdowns = document.getElementsByClassName("dropdown-content");
        var i;
        for (i = 0; i < dropdowns.length; i++) {
            var openDropdown = dropdowns[i];
            if (openDropdown.classList.contains('show')) {
                openDropdown.classList.remove('show');
            }
        }
    }
}

window.onclick = function (event) {
    if (!event.target.matches('.collapsible')) {
        var collapsible = document.getElementsByClassName("content");
        var i;
        for (i = 0; i < collapsible.length; i++) {
            var openColl = collapsible[i];
            if (openColl.classList.contains('show')) 
            {
                openColl.classList.remove('show');
            }
        }
    }
}