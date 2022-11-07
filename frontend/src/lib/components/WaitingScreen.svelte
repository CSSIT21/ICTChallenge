<script lang="ts">
	import waitingImage from '../../assets/images/waiting-image.gif'
</script>

<div
	class="h-[214px] w-full mt-[245px] flex flex-col justify-center items-center"
>
	<div class="w-[145px] h-[147px]">
		<img src={waitingImage} alt="waitingImage" />
	</div>
	<div class="h-[45px] mt-3.5 text-center text-white space-y-5">
		<h1 class="font-semibold">Please waiting for your turn</h1>
		<p class="text-xs">Your turn will come soon! It shouldn’t be long</p>
	</div>
</div>
