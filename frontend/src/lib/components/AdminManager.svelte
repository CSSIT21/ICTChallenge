<script lang="ts">
	import crown from '../../assets/images/icons-medieval-crown.png'
	import AdminPanel from './AdminPanel.svelte'
	import TeamAction from './TeamAction.svelte'
	import TeamCard from './TeamCard.svelte'

	const teams = [1, 2, 3, 4, 5, 6]
</script>

<div
	class="w-full h-full flex items-stretch "
	style="background-color: rgb(248, 251, 253)"
>
	<div
		class="w-[20vw] h-full bg-white flex flex-col"
		style="box-shadow: rgb(232 232 232 / 50%) 0px -2px 2px; border-right: 1px solid rgba(217, 217, 217,.4)"
	>
		<div
			class="h-[64px] flex items-center justify-center font-bold text-lg gap-2"
			style="border-bottom: 1px solid rgb(217, 217, 217); "
		>
			<img width={32} height={32} src={crown} alt="icons-crown" /> League Manager
		</div>

		<div class="flex-1 flex flex-col">
			<TeamCard id={1} label="Team 1" name="Benchamarachuthit" />
			<TeamCard id={2} label="Team 1" name="Benchamarachuthit" />
			<TeamCard id={3} label="Team 1" name="Benchamarachuthit" />
			<TeamCard id={4} label="Team 1" name="Benchamarachuthit" />
			<TeamCard id={5} label="Team 1" name="Benchamarachuthit" />
			<TeamCard id={6} label="Team 1" name="Benchamarachuthit" />
		</div>
	</div>
	<div class="flex-1 h-full flex flex-col">
		<div
			class="flex-1 w-[calc(80vw-6rem)]"
			style="overflow-y: hidden; overflow-x: scroll;"
		>
			<div class="flex flex-row flex-nowrap">
				{#each [1, 2, 3, 4, 5, 6, 7, 8, 9, 10] as round}
					<AdminPanel {round} />
				{/each}
			</div>
		</div>
	</div>
	<div
		class="w-[6rem] h-full bg-white flex flex-col"
		style="box-shadow: rgb(232 232 232 / 50%) 0px -2px 2px;  border-left: 1px solid rgba(217, 217, 217,.4)"
	>
		<div
			class="h-[64px] flex items-center justify-center font-bold text-md gap-2"
			style="border-bottom: 1px solid rgb(217, 217, 217); "
		>
			Round 999
		</div>
		<div class="flex-1 flex flex-col">
			<TeamAction />
			<TeamAction />
			<TeamAction />
			<TeamAction />
			<TeamAction />
			<TeamAction />
		</div>
	</div>
</div>
