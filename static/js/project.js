$(function(){   
	$('.am-btn.am-btn-primary').click(function(){
		var name=$('#name').val();
		var ipAddress=$('#ipAddress').val();
		var username=$('#username').val();
		var password=$('#password').val();
		var path=$('#path').val();
		var port=$('#port').val();
		var remark=$('#remark').val();
			if(name==''){
				alert('����������Ϊ��');
				name.focus();
			}else if(ipAddress==''){
				alert('IP��ַ����Ϊ��');
			}else if(username==''){
				alert('�û�������Ϊ��');
				username.focus();
			}else if(password==''){
				alert('���벻��Ϊ��');
				password.focus();
			}else if(path==''){
				alert('����·������Ϊ��');
				path.focus();
			}else if(port==''){
				alert('�˿ںŲ���Ϊ��');
				port.focus();
			}else if(remark==''){
				alert('��ע����Ϊ��');
				remark.focus();
			}else{
				$.ajax({
					cache: true,
					type: "POST",
					url:"/project/save",
					data:$('#project').serialize(),
					async: false,
					error: function(request) {
						alert("Connection error");
					},
					success: function(data) {
						$("#commonLayout_appcreshi").parent().html(data);
					}
				});
			}
   })

});