import type { Component } from 'solid-js';

const App: Component = () => {
  return (
    <div class='container mx-auto sm:px-4'>
      <div class='flex flex-wrap'>
        <div class='relative flex-grow max-w-full flex-1 px-4'>
          <h1 class='mt-5'>Test microservices</h1>

          <div
            id='output'
            class='mt-5'
            style='outline: 1px solid silver; padding: 2em;'
          >
            <span class='text-gray-700'>Output shows here...</span>
          </div>
        </div>
      </div>
      <div class='flex flex-wrap '>
        <div class='relative flex-grow max-w-full flex-1 px-4'>
          <h4 class='mt-5'>Sent</h4>
          <div class='mt-1' style='outline: 1px solid silver; padding: 2em;'>
            <pre id='payload'>
              <span class='text-gray-700'>Nothing sent yet...</span>
            </pre>
          </div>
        </div>
        <div class='relative flex-grow max-w-full flex-1 px-4'>
          <h4 class='mt-5'>Received</h4>
          <div class='mt-1' style='outline: 1px solid silver; padding: 2em;'>
            <pre id='received'>
              <span class='text-gray-700'>Nothing received yet...</span>
            </pre>
          </div>
        </div>
      </div>
    </div>
  );
};

export default App;
