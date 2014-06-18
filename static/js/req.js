/* req.js - Firewall Request Form Javascript */

function add_firewall_properties_form ()
{
	var table = document.getElementById('firewall_properties');
	var lastRow = table.rows.length - 1;

	var row = table.insertRow(lastRow);
	var cell = row.insertCell(0);
	var newrow = '<input class="form-control input-sm" id="source_ips_'+lastRow+'" name="source_ips_'+lastRow+'" type="text" value="">';

	lastRow = lastRow - 1;
	/*jshint multistr: true */
	var new_firewall_properties_row = '<td> \
		<select class="form-control" id="rules['+lastRow+'].source_zone" name="rules['+lastRow+'].source_zone"><option value="nonprod">Non-Prod</option><option value="prod">Prod</option><option value="pci">PCI</option><option value="dev">Dev</option></select> \
		</td> \
		<td><input class="form-control input-sm" id="rules['+lastRow+'].source_ips" name="rules['+lastRow+'].source_ips" type="text" value=""></td> \
		<td> \
		<select class="form-control" id="rules['+lastRow+'].network_protocol" name="rules['+lastRow+'].network_protocol"><option value="tcp">TCP</option><option value="udp">UDP</option><option value="tcpudp">TCP/UDP</option><option value="imcp">ICMP</option></select> \
		</td> \
		<td> \
		<select class="form-control" id="rules['+lastRow+'].dest_zone" name="rules['+lastRow+'].dest_zone"><option value="nonprod">Non-Prod</option><option value="prod">Prod</option><option value="pci">PCI</option><option value="dev">Dev</option></select> \
		</td> \
		<td> \
		<input class="form-control input-sm" id="rules['+lastRow+'].dest_ips" name="rules['+lastRow+'].dest_ips" type="text" value=""> \
		</td> \
		<td> \
		<input class="form-control input-sm" id="rules['+lastRow+'].dest_ports" name="rules['+lastRow+'].dest_ports" type="text" value=""> \
		</td> \
		<td> \
		<input class="form-control input-sm" id="rules['+lastRow+'].app_name" name="rules['+lastRow+'].app_name" type="text" value=""> \
		</td> \
		<td> \
		<input class="form-control input-sm" id="rules['+lastRow+'].server_loc" name="rules['+lastRow+'].server_loc" type="text" value=""> \
		</td> \
		<td> \
		<select class="form-control" id="rules['+lastRow+'].data_type" name="rules['+lastRow+'].data_type"><option value="pci">PCI Data</option><option value="nonpci">Non-PCI Data</option></select> \
		</td>';

	row.innerHTML=new_firewall_properties_row;
}

function alert_callback(data){
	alert(data);
}

function redirect_callback(data){
	window.location.assign(data);
}

function submit_request(){
	var jsondata = '';
	form2json(jsondata);
	
}

function form2json (jsondata){
	/*var firewall_form = document.forms["firewall_form"];*/
	var firewall_form = document.getElementById('firewall_form');
	/* var requestor = firewall_form.elements["requestor"]; */
	requestor = document.getElementById("requestor").value;
	/*jsondata = JSON.stringify(requestor);*/
	

	/* WORKING */
	/* jsondata = new form2object(firewall_form); */
	//jsondata = form2object('');
	jsondata = form2js('firewall_form', '.', false);


	/* jsondata = JSON.stringify(firewall_form.serializeObject()); */
	/* ajaxJSONPost(jsondata, alert_callback); */
	ajaxJSONPost(jsondata, redirect_callback);
}

function ajaxJSONPost(jsondata, callback){
	var xhr = new XMLHttpRequest();
	xhr.open("POST", "/req");
	xhr.setRequestHeader('Content-Type', 'application/json');
	var xsrf = document.cookie.split('=')[1];
	xhr.setRequestHeader('X-XSRF-TOKEN', xsrf);
	xhr.onreadystatechange = function () {
		if (xhr.readyState == 4 && xhr.status == 200) {
			callback(xhr.responseText);
		}
	};
	xhr.send(JSON.stringify(jsondata));
}
