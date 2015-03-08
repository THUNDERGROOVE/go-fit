
function PickerChange() {
	newSet = GetArrayFromMapping(select.value);
	var sel = document.getElementById("suit-picker")
	sel.options.length = 0;
	for (var i = newSet.length - 1; i >= 0; i--) {
		sel.options.add(new Option(newSet[i].text, newSet[i].value))
	};
}
function GetArrayFromMapping(val) {
	var out = [];
	$.each(TypeMapping, function(k, v){
		if (k == val) {
			out = v
			return
		}
	});
	return out
}

var select = document.getElementById("type-picker");
select.options.length = 0;
for (var i = suitTypes.length - 1; i >= 0; i--) {
	select.options.add(new Option(suitTypes[i].text,suitTypes[i].text))
};
PickerChange()
