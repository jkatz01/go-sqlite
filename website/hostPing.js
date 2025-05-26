function getTimestamp() {
	let d = new Date();
	const monthNames = ["January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"];
	var datestring = monthNames[(d.getMonth())] + " " + d.getDate() + " " + d.getHours() + ":" + d.getMinutes() + ":" + d.getSeconds() + "." + d.getMilliseconds();
	return datestring;
}

function someText() {
	
	curr_time = getTimestamp();
	document.getElementById("demo").innerHTML = "<p>Date: " + curr_time+ "</p>";
	
}

fetch('http://localhost:1337/send_ping', {
  method: 'POST',
  headers: {
    'Content-Type': 'application/json'
  },
  body: JSON.stringify({host_name: 'safafrvfsd', last_ping: 'May 25 20:34:28.818'})
}).then(res => res.json())
  .then(res => console.log(res));