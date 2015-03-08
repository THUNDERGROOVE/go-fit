var Scouts = [
	{value: 363956, text: "Scout A-I"},
	{value: 364025, text: "Scout A/1-Series"},
	{value: 364029, text: "Scout ak.0"},
	{value: 363955, text: "Scout C-I"},
	{value: 364024, text: "Scout C/1-Series"},
	{value: 364027, text: "Scout ck.0"},
	{value: 352587, text: "Scout G-I"},
	{value: 353759, text: "Scout G/1-Series"},
	{value: 353760, text: "Scout gk.0"},
	{value: 363957, text: "Scout M-I"},
	{value: 364026, text: "Scout M/1-Series"},
	{value: 364028, text: "Scout mk.0"},
	{value: 364178, text: "'Black Eagle' Scout G/1-Series"},
	{value: 352939, text: "Scout Type-II"},
	{value: 352940, text: "Scout B-Series"},
	{value: 352942, text: "Scout vk.1"},
	{value: 367742, text: "Scotsman's Modified Scout gk.0"},
	{value: 367900, text: "Scotsman's Modified Scout gk.0 (Master)"}
];

var Assaults = [
	{value: 352934, text: "Assault B-Series"},
	{value: 352937, text: "Assault vk.1"},
	{value: 352938, text: "Assault Type-II"},
	{value: 363349, text: "Balac's Modified Assault ck.0"},
	{value: 353763, text: "Assault C/1-Series"},
	{value: 353764, text: "Assault ck.0"},
	{value: 367738, text: "Frame's Modified Assault ck.0"},
	{value: 363936, text: "Assault M-I"},
	{value: 364018, text: "Assault M/1-Series"},
	{value: 364019, text: "Assault G/1-Series"},
	{value: 364022, text: "Assault ak.0"},
	{value: 364023, text: "Assault gk.0"},
	{value: 364319, text: "Senior Recruiter Assault C-I"},
	{value: 364322, text: "Master Recruiter Assault C-II"},
	{value: 363935, text: "Assault G-I"},
	{value: 364020, text: "Assault A/1-Series"},
	{value: 364021, text: "Assault mk.0"},
	{value: 363934, text: "Assault A-I"},
	{value: 351071, text: "Assault C-I"},
	{value: 367760, text: "Rattati's Modified Assault gk.0"},
	{value: 367897, text: "Rattati's Modified Assault gk.0 (Master)"},
	{value: 367899, text: "Frame's Modified Assault ck.0 (Master)"}
];

var Logistics = [
	{value: 352588, text: "Logistics M-I"},
	{value: 356572, text: "Logistics vk.1"},
	{value: 353768, text: "Logistics M/1-Series"},
	{value: 353769, text: "Logistics mk.0"},
	{value: 363982, text: "Logistics C-I"},
	{value: 363983, text: "Logistics G-I"},
	{value: 363984, text: "Logistics A-I"},
	{value: 355516, text: "Logistics Type-II"},
	{value: 364030, text: "Logistics C/1-Series"},
	{value: 364031, text: "Logistics G/1-Series"},
	{value: 364032, text: "Logistics A/1-Series"},
	{value: 364033, text: "Logistics ck.0"},
	{value: 364034, text: "Logistics gk.0"},
	{value: 364035, text: "Logistics ak.0"},
	{value: 356571, text: "Logistics B-Series"},
	{value: 367751, text: "Logibro's Modified Logistics mk.0"},
	{value: 367896, text: "Logibro's Modified Logistics mk.0 (Master)"}
];
var Sentinel = [
	{value: 352415, text: "Sentinel A-I"},
	{value: 367716, text: "Archduke's Modified Sentinel mk.0"},
	{value: 353766, text: "Sentinel A/1-Series"},
	{value: 353767, text: "Sentinel ak.0"},
	{value: 364009, text: "Sentinel C-I"},
	{value: 364010, text: "Sentinel G-I"},
	{value: 364036, text: "Sentinel M/1-Series"},
	{value: 364037, text: "Sentinel C/1-Series"},
	{value: 364038, text: "Sentinel G/1-Series"},
	{value: 364039, text: "Sentinel mk.0"},
	{value: 364040, text: "Sentinel ck.0"},
	{value: 364041, text: "Sentinel gk.0"},
	{value: 364011, text: "Sentinel M-I"},
	{value: 367895, text: "Archduke's Modified Sentinel mk.0 (Master)"}

];
var Commandos = [
	{value: 356401, text: "Commando A-I"},
	{value: 365262, text: "Commando A/1-Series"},
	{value: 365263, text: "Commando ak.0"},
	{value: 366698, text: "Commando C-I"},
	{value: 366699, text: "Commando G-I"},
	{value: 366700, text: "Commando M-I"},
	{value: 366710, text: "Commando C/1-Series"},
	{value: 366711, text: "Commando ck.0"},
	{value: 366712, text: "Commando G/1-Series"},
	{value: 366713, text: "Commando gk.0"},
	{value: 366714, text: "Commando M/1-Series"},
	{value: 366715, text: "Commando mk.0"},
	{value: 367885, text: "Storm Raider's Modified Commando ak.0"},
	{value: 367898, text: "Storm Raider's Modified Commando ak.0 (Master)"}
];

var TypeMapping = {
	"Assaults": Assaults,
	"Scouts":Scouts,
	"Logistics":Logistics,
	"Sentinel":Sentinel,
	"Commandos":Commandos,
}

var suitTypes = [
	{value:Scouts, text: "Scouts"},
	{value:Assaults, text:"Assaults"},
	{value:Logistics, text:"Logistics"},
	{value:Sentinel, text:"Sentinel"},
	{value:Commandos, text:"Commandos"}
];