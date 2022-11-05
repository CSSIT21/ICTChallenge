<script lang="ts">
	import { onMount } from 'svelte'
	import type { Topic } from 'src/types/question'

	import QuestionCard from '../../lib/components/QuestionCard.svelte'
	import FlippedCard from './FlippedCard.svelte'

	import iconsTopic1 from '../../assets/images/icons-topic1.png'
	import iconsTopic2 from '../../assets/images/icons-topic2.png'
	import iconsTopic3 from '../../assets/images/icons-topic3.png'
	import iconsTopic4 from '../../assets/images/icons-topic4.png'

	export let question: Topic
	export let selectedId: number
	export let openCard: (cardCol: number, cardIndex: number) => void

	let icon: string = ''

	onMount(() => {
		if (selectedId == 0) {
			icon = iconsTopic1
		} else if (selectedId == 1) {
			icon = iconsTopic2
		} else if (selectedId == 2) {
			icon = iconsTopic3
		} else if (selectedId == 3) {
			icon = iconsTopic4
		}
	})
</script>

<div class="h-[390px] flex flex-col justify-between items-center">
	{#each question.cards as card, i (i)}
		{#if card.opened == true}
			<div class="h-[224px] w-[124px]">
				<QuestionCard
					img={icon}
					score={card.score}
					textColor={'text-color-' + (selectedId + 1)}
					{openCard}
					cardCol={selectedId}
					cardIndex={i}
				/>
			</div>
		{:else}
			<FlippedCard />
		{/if}
	{/each}
</div>
