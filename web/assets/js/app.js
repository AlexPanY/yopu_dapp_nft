// import {default as Web3} from 'web3';

var BASE_API_URL = "http://127.0.0.1:8832"

var API_RESP_SUCCESS = "Success";

window.YptNFT = {
	Save: function() {
		var self = this;


		var reader;
		$("#item_image").change(event=>{
			const file = event.target.files[0];
			reader = new window.FileReader();
			reader.readAsArrayBuffer(file);
			console.log(reader);
			// uploadImageToIPFS(reader);
		});

    	$("#btn-save-nft").click(function(event){
    		var data = {
    			"address": "0x7ba177f657Bf88A6D7Fa4Fffb9Ae9D6dA7bFeEC5",
    			"name": $('#item_name').val(),
    			"description": $('#item_desc').val(),
    			"image": $('#item_image').val()
    		};
    		data.name = $.trim(data.name);
    		if(data.name == "" || data.name == null || data.name == undefined) {
    			$('#item-form').addClass("was-validated");
    			return
    		}
    		data.description = $.trim(data.description);
    		if(data.description == "" || data.description == null || data.description == undefined) {
    			$('#item-form').addClass("was-validated");
    			return
    		}
    		data.image = $.trim(data.image);
    		if(data.image == "" || data.image == null || data.image == undefined) {
    			$('#item-form').addClass("was-validated");
    			return
    		}

    		// console.log(data);
    		post(BASE_API_URL+"/v1/asset/create", data, function(resp){
    			if(resp.code == API_RESP_SUCCESS) {
    				SuccessTips("Created successfully!");
    			}
    		})   
    	});
    },
    AllTokens: function() {
    	var data = {
    		query: "all"
    	};
    	post(BASE_API_URL+"/v1/asset/create", data, function(resp){
    		if(resp.code == API_RESP_SUCCESS) {
    			SuccessTips("Created successfully!");
    		}
    	}) 
    	console.log("all");
    },
    AccountTokens: function() {

    }
};


//SuccessTips
function SuccessTips(msg) {
	var self = $("#item-tips");
	self.fadeIn("fast", function() {
		self.text(msg);
		self.addClass("alert-primary");
	});
}


//post
function post(url, data, callback) {
	$.ajax({
		type: 'POST',
		contentType: "application/json",
		dataType: "json",
		url: url,
		data: JSON.stringify(data),
		success: callback,
	});
}

window.addEventListener('load', function(){
    // if(typeof web3 !== undefined){
    //     window.web3 = new Web3(web3.currentProvider);
    // } else{
    //     window.web3 = new Web3(new Web3.providers.HttpProvider("http://localhost:8545"));
    // }
});