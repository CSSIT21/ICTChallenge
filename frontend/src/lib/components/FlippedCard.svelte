<script lang="ts">
	import CheckCircleIcon from '../../assets/images/check_circle_FILL0_wght400_GRAD0_opsz48.svg'
	export let widthCard: string
	export let heightCard: string
	export let iconSize: string
</script>

<div
	class="w-[{widthCard}] h-[{heightCard}] rounded-2xl bgColor flex justify-center items-center"
>
	<img
		src={CheckCircleIcon}
		alt="icon"
		class="w-[{iconSize}] h-[{iconSize}]"
	/>
</div>

<style>
	.bgColor {
		background-color: rgba(255, 255, 255, 0.2);
	}
</style>
