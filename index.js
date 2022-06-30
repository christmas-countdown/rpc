const { Client } = require('discord-rpc');
const rpc = new Client({ transport: 'ipc' });
const christmas = require('@eartharoid/christmas');

const setActivity = () => {
	const sleeps = christmas.getSleeps();
	rpc.setActivity({
		buttons: [
			{
				label: 'Watch the live countdown',
				url: 'https://christmascountdown.live',
			},
			{
				label: 'Add the bot',
				url: 'https://christmascountdown.live/discord',
			},
		],
		details: `${sleeps} sleeps left`,
		largeImageKey: 'santa',
		state: 'until Christmas',
	});
	console.log('[%s] Updated presence (%d sleeps left)', new Date(), sleeps);
};

rpc.on('ready', () => {
	setActivity();
	setInterval(() => {
		setActivity();
	}, 60 * 60e3); // =1h, must be at least 15s (15000ms)
});

rpc.login({ clientId: '509851616216875019' }).catch(console.error);