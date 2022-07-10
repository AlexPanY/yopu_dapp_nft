// import {default as Web3} from 'web3';

var BASE_API_URL = "http://127.0.0.1:8832"

var API_RESP_SUCCESS = "Success";

var DEFAULT_ACCOUNT = {
    "Address": "0x48BFe0929c222D4584a9fb2ff988aCCBacdb96E1",
    "PrivateKey": "6cd92297ecb7da8361b225630d975df07993d788e667294e0cb7dfbfb05f39e5"
};

window.YptNFT = {
	Save: function() {
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
    			"Address": $('#account_address').val(),
				"PrivateKey": $('#account_privatekey').val(),
    			"Name": $('#item_name').val(),
    			"Description": $('#item_desc').val(),
    			"Image": $('#item_image').val()
    		};
			data.Address = $.trim(data.Address);
    		if(data.Address == "" || data.Address == null || data.Address == undefined) {
    			$('#item-form').addClass("was-validated");
    			return
    		}
			data.PrivateKey = $.trim(data.PrivateKey);
    		if(data.PrivateKey == "" || data.PrivateKey == null || data.PrivateKey == undefined) {
    			$('#item-form').addClass("was-validated");
    			return
    		}
    		data.Name = $.trim(data.Name);
    		if(data.Name == "" || data.Name == null || data.Name == undefined) {
    			$('#item-form').addClass("was-validated");
    			return
    		}
    		data.Description = $.trim(data.Description);
    		if(data.Description == "" || data.Description == null || data.Description == undefined) {
    			$('#item-form').addClass("was-validated");
    			return
    		}
    		data.Image = $.trim(data.Image);
    		if(data.Image == "" || data.Image == null || data.Image == undefined) {
    			$('#item-form').addClass("was-validated");
    			return
    		}

    		// console.log(data);
    		post(BASE_API_URL+"/v1/token/create", data, function(resp){
    			if(resp.code == API_RESP_SUCCESS) {
    				SuccessTips($("#item-tips"), "Created successfully!");
    			}
    		})   
    	});
    },
    Detail: function() {
    	var tokenID = 1;
    	var data = {
    		TokenID: tokenID,
    	};
    	post(BASE_API_URL+"/v1/token/detail", data, function(resp){
    		if(resp.code == API_RESP_SUCCESS) {
    		}
    	})
    },
    List: function() {
    	var data = {
    		query: "all"
    	};
    	post(BASE_API_URL+"/v1/token/list", data, function(resp){
    		if(resp.code == API_RESP_SUCCESS) {
    		
    		}
    	}) 
    	console.log("all");
    },
    AccountTokens: function(data) {
    	post(BASE_API_URL+"/v1/account/assets", data, function(resp){
    		if(resp.code == API_RESP_SUCCESS) {
				renderNFTCared(resp.data)
    		}
    	})
    },
	SetTokenPrice: function() {
		$('.btn-query').click(function(event){
			DEFAULT_ACCOUNT = {
				"Address": $('#q-account-address').val(),
				"PrivateKey": $('#q-account-privatekey').val(),
			}
			// console.log(DEFAULT_ACCOUNT);
			YptNFT.AccountTokens(DEFAULT_ACCOUNT);
		});
		$('.nft-tokens-list').on("click", ".btn-set-price", function(event){
			var tokenID = $(this).attr("token-id");
			$('#nft-token-id').val(tokenID);
			$('#set-token-modal').modal('show');
		});

		$('.nft-tokens-list').on("click", ".btn-buy-token", function(event){
			var tokenID = $(this).attr("token-id");
			$('#nft-token-id').val(tokenID);
			$('#buy-token-modal').modal('show');
		});
		
		$('.btn-buy').click(function(event){
			var data = {
    			"BuyerAddress": $('#account_address').val(),
				"BuyerPrivatekey": $('#account_privatekey').val(),
				"TokenID": parseInt($('#nft-token-id').val()),
			}
			data.BuyerAddress = $.trim(data.BuyerAddress);
    		if(data.BuyerAddress == "" || data.BuyerAddress == null || data.BuyerAddress == undefined) {
    			$('#item-form').addClass("was-validated");
    			return
    		}
			data.BuyerPrivatekey = $.trim(data.BuyerPrivatekey);
    		if(data.BuyerPrivatekey == "" || data.BuyerPrivatekey == null || data.BuyerPrivatekey == undefined) {
    			$('#item-form').addClass("was-validated");
    			return
    		}

			post(BASE_API_URL+"/v1/token/buy", data, function(resp){
				if(resp.code == API_RESP_SUCCESS) {
					location.reload();
				}
			}) 
			console.log(data);
		});

		$('.btn-save-price').click(function(event){
			var data = {
				"Address": DEFAULT_ACCOUNT.Address,
				"PrivateKey": DEFAULT_ACCOUNT.PrivateKey,
				"TokenID": parseInt($('#nft-token-id').val()),
				"Price": $('.nft-sale-price').val(),
			}
			data.Price = $.trim(data.Price);
    		if(data.Price == "" || data.Price == null || data.Price == undefined) {
    			$('#item-form').addClass("was-validated");
    			return
    		}

			if (data.Price <=0) {
				$('#item-form').addClass("was-validated");
    			return
			}

			post(BASE_API_URL+"/v1/token/sale", data, function(resp){
				if(resp.code == API_RESP_SUCCESS) {
					location.reload();
				}
			}) 

			console.log(data.Price);
		});
	}
};

function renderNFTCared(data) {
	var templates = ""
	for (let i = 0; i < data.length; i++) {
		
		var buyTemlate = '<div class="card-footer text-muted">'
						+ '<a href="#" class="btn btn-primary btn-buy-token" token-id='+ i +'">Buy</a>'
						+ '</div>'
		if (data[i].Price == 0) {
			buyTemlate = '<div class="card-footer text-muted">'
						+ '<a type="button" class="btn btn-danger btn-set-price" token-id='+ i +'>Set Price</a>'
						+ '</div>'
		}
		

		templates += '<div class="col-md-2"><div class="card">'
		+ '<img src="' + data[i].Image + '" class="card-img-top" height="200px;">'
		+ '<div class="card-body">'
		+ '<p font-weight="600"><a href="#">' + data[i].Name +' #' + i + '</a></p>'
		+ '<div class="card-text"><img src="./assets/images/6f8e2979d428180222796ff4a33ab929.svg" size="12" height="22px"> ' + data[i].Price + 'ETH</div>'
		+ '</div>'
		+ buyTemlate
		+ '</div>'
		+ '</div>'
	}
	$('.nft-tokens-list').html(templates)
}


//SuccessTips
function SuccessTips(elem, msg) {
	elem.fadeIn("fast", function() {
		elem.text(msg);
		elem.addClass("alert-primary");
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
    // if(typeof web3 !== undefined){
    //     window.web3 = new Web3(web3.currentProvider);
    // } else{
    //     window.web3 = new Web3(new Web3.providers.HttpProvider("http://localhost:8545"));
    // }