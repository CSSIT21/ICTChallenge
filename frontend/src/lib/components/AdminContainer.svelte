<script lang="ts">
	import iconsCrystal from '../../assets/images/icons-crystal.png'
	import icons1st from '../../assets/images/icons-1stmedal-64.png'
	import crown from '../../assets/images/icons-medieval-crown.png'
	import up from '../../assets/images/back-icon.png'
	let show = -1
	let count = 0

	function setShow(n: number) {
		return () => {
			if (show === n) {
				show = -1
				return
			}
			show = n
		}
	}

	function increase() {
		count++
	}

	function decrease() {
		count--
	}
</script>

<div class="flex flex-col p-4 gap-6 h-full w-full">
	<div class="flex items-center gap-1 fixed">
		<h1 class="text-7xl font-bold text-white flex-1">Admin Console</h1>
		<div class="flex gap-4">
			<button class="rounded shadow" />
		</div>
	</div>
	<div class="flex-1 flex items-center justify-center gap-4">
		<button
			class="{show === 0
				? 'opacity-100'
				: 'opacity-70'} transition-all flex gap-2 font-semibold text-2xl items-center justify-between rounded-2xl shadow cursor-pointer p-6 text-white bg-slate-800 hover:bg-slate-900 active:bg-slate-900"
			on:click={setShow(0)}
		>
			<img width={32} height={32} src={crown} alt="icons-crown" />
			Show Preview
		</button>
		<button
			class="{show === 1
				? 'opacity-100'
				: 'opacity-70'} transition-all flex gap-2 font-semibold text-2xl items-center justify-between rounded-2xl shadow cursor-pointer p-6 text-white bg-slate-800 hover:bg-slate-900 active:bg-slate-900"
			on:click={setShow(1)}
		>
			<img
				width={32}
				height={32}
				src={iconsCrystal}
				alt="icons-crystal"
			/> Show Leaderboard
		</button>
		<button
			class="{show === 2
				? 'opacity-100'
				: 'opacity-70'} transition-all flex gap-2 font-semibold text-2xl items-center justify-between rounded-2xl shadow cursor-pointer p-6 text-white bg-slate-800 hover:bg-slate-900 active:bg-slate-900"
			on:click={setShow(2)}
		>
			<img width={32} height={32} src={icons1st} alt="icons-medal" /> Show
			Podium
		</button>

		<div class="w-2 h-12 bg-[rgba(255,255,255,.3)] rounded-full mx-5" />

		<button
			class="transition-all flex gap-2 font-semibold text-2xl items-center justify-between rounded-2xl shadow cursor-pointer p-5 text-white border-4 border- border-white hover:bg-[rgba(255,255,255,.1)] active:bg-[rgba(255,255,255,.2)]"
			on:click={increase}
		>
			<img
				width={16}
				height={16}
				src={up}
				class="rotate-90"
				alt="icons-medal"
			/> Reveal
		</button>
		<button
			class="transition-all flex gap-2 font-semibold text-2xl items-center justify-between rounded-2xl shadow cursor-pointer p-5 text-white border-4 border- border-white hover:bg-[rgba(255,255,255,.1)] active:bg-[rgba(255,255,255,.2)]"
			on:click={decrease}
		>
			<img
				width={16}
				height={16}
				src={up}
				class="-rotate-90"
				alt="icons-medal"
			/> Unreveal
		</button>
	</div>
</div>
