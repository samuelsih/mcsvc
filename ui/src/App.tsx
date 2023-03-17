import { Component, createSignal, Show } from 'solid-js';
import { createStore } from 'solid-js/store';

type Output = {
  from: From;
  result: any;
};

enum From {
  INITIAL = 'INITIAL',
  BROKER = 'BROKER',
  ERROR = 'ERROR',
}

const [output, setOutput] = createStore<Output>({
  from: From.INITIAL,
  result: 'Output Shows Here',
});
const [payload, setPayload] = createSignal(`Nothing sent yet...`);
const [received, setReceived] = createSignal(`Nothing received yet...`);

const testBroker = async () => {
  try {
    const body = {
      method: 'POST',
    };

    const response = await fetch('http://localhost:8080', body);
    const data = await response.json();

    setPayload(`Empty post request`);
    setReceived(JSON.stringify(data, undefined, 4));
    setOutput({ from: From.BROKER, result: data.message });

    if (data.error) {
      throw new Error(data.message);
    }
  } catch (error) {
    console.error(error);
    setOutput({ from: From.ERROR, result: error });
  }
};

const App: Component = () => {
  return (
    <div class='container mx-auto sm:px-4'>
      <div class='flex flex-wrap'>
        <div class='relative flex-grow max-w-full flex-1 px-4 space-y-4'>
          <h1 class='mt-5'>Test microservices</h1>
          <button
            class='class="inline-block align-middle text-center select-none border font-normal whitespace-no-wrap rounded py-1 px-3 leading-normal no-underline text-gray-600 border-gray-600 hover:bg-gray-600 hover:text-white bg-white hover:bg-gray-700"'
            onClick={testBroker}
          >
            Test Broker
          </button>
          <div
            id='output'
            class='mt-5'
            style='outline: 1px solid silver; padding: 2em;'
          >
            <pre class='text-gray-700 flex flex-col'>
              <Show
                when={output.from !== From.INITIAL}
                fallback={`See output here`}
              >
                <p>
                  <strong>Response from {output.from} service</strong> : {output.result}
                </p>
              </Show>
            </pre>
          </div>
        </div>
      </div>
      <div class='flex flex-wrap '>
        <div class='relative flex-grow max-w-full flex-1 px-4'>
          <h4 class='mt-5'>Sent</h4>
          <div class='mt-1' style='outline: 1px solid silver; padding: 2em;'>
            <pre id='payload'>
              <span class='text-gray-700'>{payload()}</span>
            </pre>
          </div>
        </div>
        <div class='relative flex-grow max-w-full flex-1 px-4'>
          <h4 class='mt-5'>Received</h4>
          <div class='mt-1' style='outline: 1px solid silver; padding: 2em;'>
            <pre id='received'>
              <span class='text-gray-700'>{received()}</span>
            </pre>
          </div>
        </div>
      </div>
    </div>
  );
};

export default App;
