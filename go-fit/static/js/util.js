String.prototype.capitalize = function() {
	return this.charAt(0).toUpperCase() + this.slice(1);
}

function HTTPGet(url) {
	var out = "";
	$.ajax({
	url: url,
	type: "GET",
	async: false,
	success: function(data) {
		out = data
		}
	});
	return out
}

function SaveFit(fitID) {
	var SDEType = Fit["SDEType"];
	Fit["SDEType"] = undefined;
	$.ajax({
		type:"POST",
		url:"/fit/save/"+fitID,
		data: JSON.stringify(Fit),
		success: function(data) {
			if (data== "success") {
				$("#alert-success-text").html("Fit Saved!")
				$("#alert-success").show()
			} else {
				$("#alert-error-text").html("An error occured:"+data)
				$("#alert-error").show()
			}
		}
	})
	Fit["SDEType"] = SDEType;
}

function DeleteFit(fitID) {
	$.ajax({
	type:"GET",
	url:"/fit/delete/"+fitID,
	success: function(data) {
		if (data== "success") {
			$("#alert-success-text").html("Fit deleted!")
			$("#alert-success").show()
			$("#fit-"+fitID).remove()
		} else {
			$("#alert-error-text").html("An error occured:"+data)
			$("#alert-error").show()
		}
	}})
}

function RenameFit(fitID, defaultText) {
	var n = window.prompt("New name",defaultText);
	$.ajax({
	type:"POST",
	url:"/fit/rename/"+fitID,
	data:n,
	success: function(data) {
		if (data== "success") {
			$("#alert-success-text").html("Fit Saved!")
			$("#alert-success").show()
			$("#fit-"+fitID+"-name").html(n)
		} else {
			$("#alert-error-text").html("An error occured:"+data)
			$("#alert-error").show()
		}
	}
	})
}