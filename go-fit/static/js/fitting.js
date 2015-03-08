function DrawSlots(count, className) {
	if(typeof count === 'undefined'){
		return
	};
	if (count==0) {
		return
	};
	for (var i = count-1; i >= 0; i--) {
		$("#" + className).append('\
				<div id="' + className +'-' + i + '" class="dropdown"> \
					<button id="' + className +'-' + i + '-button" class="btn btn-default dropdown-toggle" type="button" id="dropdownMenu1" data-toggle="dropdown" aria-expanded="true">\
						ItemSlot\
						<span class="caret"></span>\
					</button>\
						<ul class="dropdown-menu" role="menu" aria-labelledby="dropdownMenu1">\
						<li role="presentation"><a role="menuitem" onclick="ClearItem(\'' + className +'-' + i + '\');" tabindex="-1" href="#">Clear Slot</a></li>\
						<li role="presentation"><a role="menuitem" onclick="ChangeItem(\'' + className +'-' + i + '\');" tabindex="-1" href="#">Change Item</a></li>\
					</ul>\
				</div>')
	};
}

var Fit = [];

function LoadFitJSON(fitID) {
	var json = null;
	$.ajax({
		'async': false,
		'global': false,
		'url': "/rawfit/"+fitID,
		'dataType': "json",
		'success': function (data) {
		json = data;
	}
	});
	return json;
}

function CountSlotType(attr, slotType) {
	var c = 0;
	for (var a in attr) {
		if (a.includes("mModuleSlots.")) {
			if (a.split(".")[2] == "slotType" && attr[a] == slotType) {
				c++;
			};
		};
	};
	return c;
}

function GenerateSlots(fitID) {
	Fit = LoadFitJSON(fitID)
	attr = Fit["SDEType"]["attributes"]
	console.log(attr)
	hs = CountSlotType(attr, "IH")
	DrawSlots(hs, "high-slots")
	ls = CountSlotType(attr, "IL")
	DrawSlots(ls, "low-slots")
	wp = CountSlotType(attr, "WP")
	DrawSlots(wp, "light-slots")
	sa = CountSlotType(attr, "WS")
	DrawSlots(sa, "sidearm-slots")
	wh = CountSlotType(attr, "WH")
	DrawSlots(wh, "heavy-slots")
	gs = CountSlotType(attr, "GS")
	DrawSlots(gs, "grenade-slots")
	eq = CountSlotType(attr, "IE")
	DrawSlots(eq, "equipment-slots")
	OnChangeFit()
}

var currentSel = 0;

function ChangeItem(slotClass) {
	currentSel = slotClass
	$("#item-picker").modal({show:true});
	SetupPicker();
}

function ChooseItem(slotClass) {
	var sel = document.getElementById("module-picker");
	var slot = parseInt(slotClass.split("-")[2],10)+1;
	Fit[slotClass.split("-")[0].capitalize() + "Slot" + slot] = parseInt(sel.value,10);
	OnChangeFit();
}

function OnChangeFit() {
	for (var k in Fit) {
		if (k.includes("HighSlot")) {
			if (Fit[k]!= 0) {
				var slot = k.split("HighSlot")[1]-1
				$.ajax({
					url: "/typename/"+Fit[k],
					type: "GET",
					async: false,
					success: function(data) {
						$("#high-slots-"+slot+"-button").html(data+"<span class=\"caret\"></span>")
						console.log("Found high slot "+slot + " with typeid:" +Fit[k], " name: "+data)
					}
				});
			}
		}
		if (k.includes("LowSlot")) {
			if (Fit[k]!= 0) {
				var slot = k.split("LowSlot")[1]-1
				$.ajax({
					url: "/typename/"+Fit[k],
					type: "GET",
					async: false,
					success: function(data) {
						$("#low-slots-"+slot+"-button").html(data+"<span class=\"caret\"></span>")
						console.log("Found low slot "+slot + " with typeid:" +Fit[k], " name: "+data)
					}
				});
			}
		}
	}
}

function UpdateSlotName() {
	
}