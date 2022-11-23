<script lang="ts">
	import { onMount } from 'svelte'
	import type { Topic } from 'src/types/question'
	import { selected, current } from 'src/store/system'
	import { axios } from '../../utils/api'

	import QuestionCard from '../../lib/components/QuestionCard.svelte'
	import FlippedCard from './FlippedCard.svelte'

	import iconsTopic1 from '../../assets/images/icons-topic1.png'
	import iconsTopic2 from '../../assets/images/icons-topic2.png'
	import iconsTopic3 from '../../assets/images/icons-topic3.png'
	import iconsTopic4 from '../../assets/images/icons-topic4.png'
	import iconsTopic5 from '../../assets/images/icons-topic5.png'

	export let question: Topic

	let icon: string = ''
	const widthCard: string = '224px'
	const heightCard: string = '124px'
	const widthImg: string = '60.3px'
	const textSize: string = '27px'
	const iconSize: string = '80px'

	const handleFlipCard = async (id: number, index: number) => {
		console.log(id, index)
		axios.patch('http://ictc-int.sit.kmutt.ac.th:3000/api/am/card/pause', {},{
				headers: {
					Authorization: 'Bearer 7GRPpVk04We76CWR',
					'Content-Type': 'application/json',
				},
			})
		await axios.put(
			'http://ictc-int.sit.kmutt.ac.th:3000/api/st/open',
			JSON.stringify({
				topic_id: id,
				card_id: index,
			}),
			{
				headers: {
					Authorization: 'Bearer XrO3ole8bS83OQ3p',
					'Content-Type': 'application/json',
				},
			}
		)
		setTimeout(() => resetStatus(), 2000)
	}

	const resetStatus = () => {
		current.set(0)
		selected.set(-1)
		document.body.scrollIntoView()
	}

	onMount(() => {
		if ($selected == 0) {
			icon = iconsTopic1
		} else if ($selected == 1) {
			icon = iconsTopic2
		} else if ($selected == 2) {
			icon = iconsTopic3
		} else if ($selected == 3) {
			icon = iconsTopic4
		} else if ($selected == 4) {
			icon = iconsTopic5
		}
	})
</script>

<div class="mt-12">
	<div class="h-[480px] w-[390px] flex flex-col justify-between">
		{#each question.cards as card, i (i)}
			<div class={card.opened && 'flipped'}>
				<div class="flip-card">
					<div class="flip-card-inner flex">
						<div class="flip-card-front flex justify-center">
							<QuestionCard
								img={icon}
								score={card.score}
								textColor={'text-color-' + ($selected + 1)}
								{widthCard}
								{heightCard}
								{widthImg}
								{textSize}
								{handleFlipCard}
								cardId={card.id}
								topicId={$selected + 1}
							/>
						</div>
						<div
							class="flip-card-back flex justify-center"
							style="transform: rotateY(180deg);"
						>
							<FlippedCard {widthCard} {heightCard} {iconSize} />
						</div>
					</div>
				</div>
			</div>
		{/each}
	</div>
</div>

<style>
	.flip-card {
		cursor: pointer;
		user-select: none;
	}
	.flip-card-inner {
		width: 100%;
		height: 100%;
		text-align: center;
		transition: transform 0.6s;
		transform-style: preserve-3d;
		box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2);
	}

	.flipped .flip-card .flip-card-inner {
		transform: rotateY(180deg);
	}

	.flip-card-front,
	.flip-card-back {
		position: absolute;
		width: 100%;
		height: 100%;
		-webkit-backface-visibility: hidden;
		backface-visibility: hidden;
	}
</style>
