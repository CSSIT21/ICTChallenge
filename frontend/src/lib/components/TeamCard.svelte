<script lang="ts">
	export let id = 0
	export let label = 'Team 1'
	export let name = 'Benchamarachuthit'
</script>

<div
	class="flex flex-row gap-2 h-[72px] items-center justify-center"
	style="border-bottom: 1px solid rgba(217, 217, 217, .5)"
>
	<div class="w-[64px] flex items-center justify-center">
		<div
			class="rounded-full w-[32px] h-[32px] flex justify-center items-center font-bold"
			style="background-color: rgb(242, 242, 242); color: rgba(0,0,0,.4)"
		>
			{id}
		</div>
	</div>
	<div class="flex flex-col items-center justify-center">
		<b>{label}</b>
		<span class="ml-2">{name}</span>
	</div>
</div>
