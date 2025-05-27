function getTimestamp() {
	let d = new Date();
	const monthNames = ["January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"];
	var datestring = monthNames[(d.getMonth())] + " " + d.getDate() + " " + d.getHours() + ":" + d.getMinutes() + ":" + d.getSeconds() + "." + String(d.getMilliseconds()).padStart(3, '0');
	return datestring;
}

function someText() {
	
	curr_time = getTimestamp();
	document.getElementById("date").innerHTML = "Date: " + curr_time;
	
}

function pingRandomHost() {
	const hostNames = ["Somebody", "someone", "some dude", "buddy", "guh", "jsbrowser", "mike", "gus", "walter"];
	var curr_time = getTimestamp();
	fetch('http://localhost:1337/send_ping', {
	mode: 'no-cors',
	method: 'POST',
	headers: {
		'Content-Type': 'application/json'
	},
	body: JSON.stringify({host_name: hostNames[Math.floor(Math.random() * 9)], ping_time: curr_time})
	});
	document.getElementById("date").innerHTML = "Date: " + curr_time;
}


pingRandomHost();
someText();